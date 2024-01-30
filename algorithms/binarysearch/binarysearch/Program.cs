
class BinarySearch {
    static void Main(string[] args) {

    }

    static void Binary(int[] numbers, int num) {
        
        int len = 0;
        if(numbers.Length % 2 == 0) {
            len = numbers.Length / 2;
        }
        else {
            len = (numbers.Length - 1)/2;
        }
        
        bool trigger = true;
        while(trigger) {
            if(num == numbers[len]) {
                trigger = false;
                break;
            }
            else if (num < numbers[len/2]) {
                len /= 2;
            }
            else if (num > numbers[len/2]) {
                len += (len/2);
            }
        }
    }
}