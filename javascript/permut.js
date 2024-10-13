function permuteUnique(s) {
    const result = [];
    const used = Array(s.length).fill(false);
    const chars = s.split('').sort(); // Sort to handle duplicates

    const backtrack = (path) => {
        if (path.length === s.length) {
            result.push(path.join(''));
            return;
        }
        for (let i = 0; i < chars.length; i++) {
            if (used[i]) continue; // Skip used characters
            if (i > 0 && chars[i] === chars[i - 1] && !used[i - 1]) continue; // Skip duplicates

            used[i] = true; // Mark as used
            path.push(chars[i]);
            backtrack(path); // Recursive call
            path.pop(); // Backtrack
            used[i] = false; // Unmark as used
        }
    };

    backtrack([]);
    return result;
}

module.exports = permuteUnique;
// Example usage
if (require.main === module) {
    console.log(permuteUnique("aabb")); // Output: ["aabb", "abab", "abba", "baab", "baba", "abba"]
}
