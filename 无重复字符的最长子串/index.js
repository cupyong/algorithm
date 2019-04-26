function longOfLengthSubstring(s) {
    if (!s) return 0;
    var map = {};
    var maxLength = 0;
    var j = 0;
    for (var i = 0; i < s.length; i++) {
        var c = s[i];
        if (map[c]) {
            j = Math.max(j, map[c] + 1)
        }
        map[c] = i;
        maxLength = Math.max(maxLength, i - j + 1)
    }
    return maxLength
}

function Test() {
    var s = "abcdefasd"
    console.log(longOfLengthSubstring(s))
}

Test()