function reverse(x) {
    var y = 0;
    while (x != 0) {
        y = y*10 + x%10
        x =parseInt(x/10)
        console.log(x)
    }
    return y
}
console.log(reverse(-120))