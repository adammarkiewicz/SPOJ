using System;

namespace BFN1
{
    class Program
    {
        static void Main()
        {
            short testsCount = Convert.ToInt16(Console.ReadLine());
            string number;
            int additionsCounter;

            for (short c = 0; c < testsCount; c++)
            {
                number = Console.ReadLine();
                additionsCounter = 0;

                while (number != ReverseString(number))
                {
                    number = (Convert.ToInt16(number) + Convert.ToInt16(ReverseString(number))).ToString();
                    additionsCounter++;
                }

                Console.WriteLine(number + " " + additionsCounter);
            }
        }

        private static string ReverseString(string s)
        {
            char[] arr = s.ToCharArray();
            Array.Reverse(arr);
            return new string(arr);
        }
    }
}
