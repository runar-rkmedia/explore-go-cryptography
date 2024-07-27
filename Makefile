testkey := "deaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddead"
GO_FILES := $(shell find . -type f -iname "*.go")

test_all: test cli_test
test:
	gotestsum  ./... -- -count=1 | ,colored-logs
test_watch:
	fd | entr -rc  gotestsum  ./... -- -count=1 | ,colored-logs
cli_test:
	./cli_test.sh
./enciphered.bin: ./tiger.txt $(GO_FILES)
	go run ./shift/cmd/encipher -key $(testkey) < tiger.txt > /tmp/enciphered.bin && mv /tmp/enciphered.bin ./enciphered.bin
decode-enciphered.bin: ./enciphered.bin $(GO_FILES)
	go run ./shift/cmd/decipher -key $(testkey) < enciphered.bin
devil.ppm: ./devil.png
	magick devil.png devil.ppm
header.txt: ./devil.ppm
	head -n 3 devil.ppm > header.txt
imgdata.txt: devil.ppm
	tail -n +4 devil.ppm >imgdata.txt
devil.ppm.bin: imgdata.txt $(GO_FILES)
	go run ./shift/cmd/encipher -key 0101010101010101010101010101010101010101010101010101010101010101  < imgdata.txt > devil.ppm.bin
test.ppm: devil.ppm.bin header.txt
	cat header.txt devil.ppm.bin > test.ppm




