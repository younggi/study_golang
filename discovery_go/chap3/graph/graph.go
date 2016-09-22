package graph

import (
	"fmt"
	"io"
)

// WriteTo writes to w with adjList
func WriteTo(w io.Writer, adjList [][]int) error {
	size := len(adjList)
	if _, err := fmt.Fprintf(w, "%d", size); err != nil {
		return err
	}
	for i := 0; i < size; i++ {
		lsize := len(adjList[i])
		if _, err := fmt.Fprintf(w, "\n%d", lsize); err != nil {
			return err
		}
		for j := 0; j < lsize; j++ {
			if _, err := fmt.Fprintf(w, " %d", adjList[i][j]); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
	return nil
}

// ReadFrom reads from r to adjList
func ReadFrom(r io.Reader, adjList *[][]int) error {
	var size int
	if _, err := fmt.Fscanf(r, "%d", &size); err != nil {
		return err
	}
	*adjList = make([][]int, size)
	for i := 0; i < size; i++ {
		var lsize int
		if _, err := fmt.Fscanf(r, "\n%d", &lsize); err != nil {
			return err
		}
		(*adjList)[i] = make([]int, lsize)
		for j := 0; j < lsize; j++ {
			if _, err := fmt.Fscanf(r, " %d", &(*adjList)[i][j]); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fscanf(r, "\n"); err != nil {
		return err
	}
	return nil
}
