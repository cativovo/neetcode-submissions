class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {
            if (s.length !== t.length) {
      return false;
    }

    const sM = {};
    const tM = {};

    for (let i = 0; i < s.length; i++) {
      const sC = s[i];
      sM[sC] = (sM[sC] ?? 0) + 1;

      const tC = t[i];
      tM[tC] = (tM[tC] ?? 0) + 1;
    }

    for (const key in sM) {
      if (sM[key] !== tM[key]) {
        return false;
      }
    }

    return true;

    }
}
