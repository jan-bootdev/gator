package main

import "context"

func handlerReset(s *state, _ command) error {

	err := s.db.Reset(context.Background())
	if err != nil {
		return err
	}

	return nil

}
