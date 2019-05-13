const letterCombinations = function(digits) {
    if (!digits) return [];
    let dict = [
        0,
        1,
        ['a', 'b', 'c'],
        ['d', 'e', 'f'],
        ['g', 'h', 'i'],
        ['j', 'k', 'l'],
        ['m', 'n', 'o'],
        ['p', 'q', 'r', 's'],
        ['t', 'u', 'v'],
        ['w', 'x', 'y', 'z'],
    ];
    let ans = [''];
    for (let digit of digits) {
        let curAns = [];
        for (let letter of dict[+digit]) {
            ans.forEach((word) => curAns.push(word+letter));
        }
        ans = curAns;
        console.log(curAns)
    }
    return ans;
};
console.log(letterCombinations("345"))
