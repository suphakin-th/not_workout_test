function countSmileys(arr) {
    const smileyRegex = /^[:;][-~]?[)D]$/;
    return arr.filter(face => smileyRegex.test(face)).length;
}

module.exports = countSmileys;

if (require.main === module) {
    const arr = [":)", ";(", ";}", ":-D"];
    console.log("Count of smiley faces:", countSmileys(arr)); // Output: 2
}
