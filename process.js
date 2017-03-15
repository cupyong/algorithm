


process.on('exit', (code) => {
    console.log(`About to exit with code: ${code}`);
});

console.log(11111)
process.on('beforeExit', (code) => {
    console.log(`About to beforeExit with code: ${code}`);
});
process.argv.forEach((val, index) => {
    console.log(`${index}: ${val}`);
});