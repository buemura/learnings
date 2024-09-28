function reverseOnlyLetters(s: string): string {
  const letters: string[] = [];
  let reversedResult = "";

  for (const char of s) {
    if (char.match(/^[a-zA-Z]/)) {
      letters.push(char);
    }
  }

  for (const char of s) {
    if (char.match(/^[a-zA-Z]/)) {
      reversedResult += letters.pop();
    } else {
      reversedResult += char;
    }
  }

  return reversedResult;
}

console.log(reverseOnlyLetters("ab-cd")); // dc-ba
console.log(reverseOnlyLetters("a-b-cd")); // d-c-ba
console.log(reverseOnlyLetters("a-bC-dEf=ghIj!!")); // j-Ih-gfE=dCba!!
