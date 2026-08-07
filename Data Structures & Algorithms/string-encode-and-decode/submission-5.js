class Solution {
  delimiter = "*";

  /**
   * @param {string[]} strs
   * @returns {string}
   */
  encode(strs) {
    let r = "";

    for (const v of strs) {
      r += `${v.length}${this.delimiter}${v}`;
    }

    return r;
  }

  /**
   * @param {string} str
   * @returns {string[]}
   */
  decode(str) {
    const r = [];

    while (str !== "") {
      const i = str.indexOf(this.delimiter);
      const c = parseInt(str.substring(0, i));
      str = str.substring(i + 1);
      r.push(str.substring(0, c));
      str = str.substring(c);
    }

    return r;
  }
}
