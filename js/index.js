'use strict';
//判断一个单词是否是回文？
function checkPalindrom(str) {
    return str == str.split('').reverse().join('');
}
console.log(checkPalindrom("abccba"))

//Q2 去掉一组整型数组重复的值
// 比如输入: [1,13,24,11,11,14,1,2]
// 输出: [1,13,24,11,14,2]
// 需要去掉重复的11 和 1 这两个元素

let unique = function(arr) {
    let hashTable = {};
    let data = [];
    for(let i=0,l=arr.length;i<l;i++) {
        if(!hashTable[arr[i]]) {
            hashTable[arr[i]] = true;
            data.push(arr[i]);
        }
    }
    return data

}
console.log(unique([1,13,24,11,11,14,1,2]))

//Q3 统计一个字符串出现最多的字母
function findMaxDuplicateChar(str) {
    if(str.length == 1) {
        return str;
    }
    let charObj = {};
    for(let i=0;i<str.length;i++) {
        if(!charObj[str.charAt(i)]) {
            charObj[str.charAt(i)] = 1;
        }else{
            charObj[str.charAt(i)] += 1;
        }
    }
    let maxChar = '',
        maxValue = 1;
    for(var k in charObj) {
        if(charObj[k] >= maxValue) {
            maxChar = k;
            maxValue = charObj[k];
        }
    }
    return maxChar;

}

console.log(findMaxDuplicateChar("afjghdfraaaasdenas"))
//快排
function quickSort(arr) {

    if(arr.length<=1) {
        return arr;
    }

    let leftArr = [];
    let rightArr = [];
    let q = arr[0];
    for(let i = 1,l=arr.length; i<l; i++) {
        if(arr[i]>q) {
            rightArr.push(arr[i]);
        }else{
            leftArr.push(arr[i]);
        }
    }

    return [].concat(quickSort(leftArr),[q],quickSort(rightArr));
}

console.log(quickSort([1,13,24,11,11,14,1,2]))

//交换
function swap(a , b) {
    b = b - a;
    a = a + b;
    b = a - b;
    return [a,b];
}
//随机数
function randomString(n) {
    let str = 'abcdefghijklmnopqrstuvwxyz9876543210';
    let tmp = '',
        i = 0,
        l = str.length;
    for (i = 0; i < n; i++) {
        tmp += str.charAt(Math.floor(Math.random() * l));
    }
    return tmp;
}

console.log(randomString(4))

class Node {
    constructor(data, left, right) {
        this.data = data;
        this.left = left;
        this.right = right;
    }

}

class BinarySearchTree {

    constructor() {
        this.root = null;
    }

    insert(data) {
        let n = new Node(data, null, null);
        if (!this.root) {
            return this.root = n;
        }
        let currentNode = this.root;
        let parent = null;
        while (1) {
            parent = currentNode;
            if (data < currentNode.data) {
                currentNode = currentNode.left;
                if (currentNode === null) {
                    parent.left = n;
                    break;
                }
            } else {
                currentNode = currentNode.right;
                if (currentNode === null) {
                    parent.right = n;
                    break;
                }
            }
        }
    }

    remove(data) {
        this.root = this.removeNode(this.root, data)
    }

    removeNode(node, data) {
        if (node == null) {
            return null;
        }

        if (data == node.data) {
            // no children node
            if (node.left == null && node.right == null) {
                return null;
            }
            if (node.left == null) {
                return node.right;
            }
            if (node.right == null) {
                return node.left;
            }

            let getSmallest = function(node) {
                if(node.left === null && node.right == null) {
                    return node;
                }
                if(node.left != null) {
                    return node.left;
                }
                if(node.right !== null) {
                    return getSmallest(node.right);
                }

            }
            let temNode = getSmallest(node.right);
            node.data = temNode.data;
            node.right = this.removeNode(temNode.right,temNode.data);
            return node;

        } else if (data < node.data) {
            node.left = this.removeNode(node.left,data);
            return node;
        } else {
            node.right = this.removeNode(node.right,data);
            return node;
        }
    }

    find(data) {
        var current = this.root;
        while (current != null) {
            if (data == current.data) {
                break;
            }
            if (data < current.data) {
                current = current.left;
            } else {
                current = current.right
            }
        }
        return current.data;
    }

}

let BinarySearchTree1 = new BinarySearchTree()