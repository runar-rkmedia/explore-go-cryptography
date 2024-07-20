#!/usr/bin/env bash
testkey="deaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddead"
encipher="go run ./shift/cmd/encipher/main.go"
decipher="go run ./shift/cmd/decipher/main.go"
crack="go run ./shift/cmd/crack/main.go"

testMessages=(
  'This is 32 bytes, including EOF'
  'This is 31 long, including EOF'
  'This is a bit longer 44 long, including EOF'
)
testMessages_count=${#testMessages[@]}

for i in "${!testMessages[@]}"; do
  plaintext="${testMessages[$i]}"
  # byte_count=$(printf "%s" "$plaintext" | wc -c)

  echo "Test message $((i +1))/$testMessages_count: '${plaintext}'"
  encrypted="$(echo "${plaintext}" | $encipher -key $testkey --output-base-64)"
  echo "Enc: '$encrypted'"
  decrypted="$(echo "${encrypted}" | $decipher -key $testkey --input-base-64)"
  echo "Dec: '$decrypted'"
  if [ "${decrypted}" == "${plaintext}" ] ; then
    echo "[Encipher -> Decipher] Success!"
  else 
    echo "[Encipher -> Decipher] Failure!"
    exit 1
  fi


# echo "${encrypted}" | base64 -d | $crack -crib This 
done

