magick devil.png devil.ppm
head -n 3 devil.ppm >header.txt
tail -n +4 devil.ppm >imgdata.txt
go run ./shift/cmd/encipher -key 0101010101010101010101010101010101010101010101010101010101010101 <imgdata.txt >devil.ppm.bin
cat header.txt devil.ppm.bin >test.ppm
echo "Wrote test.ppm. Open with an image-viewer"
