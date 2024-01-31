
class BinarySearch {
    static void Main(string[] args) {
        Binary([-6,-6,-2,-1,0,1,2,3,4,5,6,7,8,8,8,9,9,10,11], 2);
        Binary([-6,-6,-2,-1,0,1,2,3,4,5,6,7,8,8,8,9,9,10,11], 5);
        Binary([-6,-6,-2,-1,0,1,2,3,4,5,6,7,8,8,8,9,9,10,11], 8);
        Binary([-6,-6,-2,-1,0,1,2,3,4,5,6,7,8,8,8,9,9,10,11], 11);
        Binary([-6,-6,-2,-1,0,1,2,3,4,5,6,7,8,8,8,9,9,10,11], -2);
        Binary([-6,-6,-2,-1,0,1,2,3,4,5,6,7,8,8,8,9,9,10,11], -6);
        Binary([-6,-6,-2,-1,0,1,2,3,4,5,6,7,8,8,8,9,9,10,11], 20);    
    }

    static void Binary(int[] numbers, int num) {
        int high = numbers.Length - 1;
        int low = 0;
        int middle;

        do {
            middle = low + (high - low) / 2;
            if(num == numbers[middle]) {
                Console.WriteLine(num);
            }
            else if(num > numbers[middle]) {
                low = middle;
            }
            else if(num < numbers[middle]) {
                high = middle;
            }
            else {
                Console.WriteLine("Your item was not found.");
            }

        } while(low < high);
    }
}