#include <iostream>
using namespace std;

bool isPalindrome(int x) {
    if (x < 0) {
        return false;
    }

    int div = 1;

    while (x >= 10 * div) {
        div *= 10;
    }

    while (x != 0) {
        int right = x % 10;
        int left = x / div;

        if (left != right) {
            return false;
        }

        x = (x % div) / 10;
        div /= 100;
    }

    return true;
}

int main() {
    // Test cases
    cout << isPalindrome(121) << endl;   // true (1)
    cout << isPalindrome(-121) << endl;  // false (0)
    cout << isPalindrome(10) << endl;    // false (0)
}
