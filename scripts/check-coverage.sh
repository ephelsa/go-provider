COVERAGE=$(cat code-coverage-results.md | grep -oi "\*\*[0-9]*%\*\*" | grep -o "[0-9]*")
MIN=$(cat code-coverage-results.md | grep -oi "\`[0-9]*%\`" | grep -o "[0-9]*")

if [ $COVERAGE -le $MIN ]; then
    echo "Failed because the current coverage ($COVERAGE%) is less than minimum ($MIN%)."
    exit 1
fi