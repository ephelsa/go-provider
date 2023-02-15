set -e

echo "" > coverage.txt

echo "Generating Coverage file."

for d in $(go list ./... | grep -v vendor); do
    go test -coverprofile=profile.out -covermode=count $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

# Modify file to the expected format
echo "Filtering the file"
sed -i "s/mode: count//g" coverage.txt
echo "mode: count" > coverage.tmp.txt
grep "\S" coverage.txt  >> coverage.tmp.txt

echo "Running coverage with cobertura"
gocover-cobertura < coverage.tmp.txt > coverage.cobertura.xml
