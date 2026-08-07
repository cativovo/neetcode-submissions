class Solution {
  /**
   * @param {string[]} strs
   * @return {string[][]}
   */
  groupAnagrams(strs) {
    const m = {};

    for (const v of strs) {
      const k = this.createKey(v);
      if (!m[k]) {
        m[k] = [v];
      } else {
        m[k].push(v);
      }
    }

    return Object.values(m);
  }

  createKey(s) {
    if (s === "") {
      return s;
    }

    const r = new Array(26).fill(0);
    const aPosition = "a".charCodeAt();

    for (const c of s) {
      r[c.charCodeAt() - aPosition]++;
    }

    return r.join("-");
  }

}
