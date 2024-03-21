package main

//import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
// Keep filling the stream infinitely with A
func (r MyReader) Read(stream []byte) (int, error) {
	n := len(stream)
	for i := 0; i < n; i++ {
		stream[i] = 'A'
	}
	return n, nil
}

// Cannot validate locally, this requires some online package

func main() {
	reader.Validate(MyReader{})
}
