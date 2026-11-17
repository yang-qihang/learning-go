package main

import "os"

func fileLen(path string) (int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	//	This is an alternatie way to achieve the same result
	//	The implementation is OK, but int64 would be a more appropriate
	//	return type then int.
	//
	//	data := make([]byte, 2048)
	//	total := 0
	//	for {
	//		count, err := f.Read(data)
	//		total += count
	//		if err != nil {
	//			if err != io.EOF {
	//				return 0, err
	//			}
	//			break
	//		}
	//	}
	//	return total, nil

	info, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
