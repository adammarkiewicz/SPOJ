using System;
using System.Text;

namespace PP0602D
{
    class Program
    {
        static void Main()
        {
            string input = Console.ReadLine().Trim();
            string[] inputStrings = input.Split(' ');

            short size = Convert.ToInt16(inputStrings[0]);
            short shift = Convert.ToInt16(inputStrings[1]);

            if (Math.Abs(shift) >= size)
            {
                shift %= size;
            }

            short[] inputArray = new short[size];
            short[] outputArray = new short[size];

            inputStrings = Console.ReadLine().Trim().Split(' ');
            for (int i = 0; i < inputStrings.Length; i++)
            {
                inputArray[i] = Convert.ToInt16(inputStrings[i]);
            }

            if (shift > 0)
            {
                for (int i = 0; i < size; i++)
                {
                    if (i + shift < size)
                        outputArray[i] = inputArray[i + shift];
                    else
                        outputArray[i] = inputArray[Math.Abs(size - shift - i)];
                }
            }
            else if (shift < 0)
            {
                for (int i = 0; i < size; i++)
                {
                    if (i + size + shift < size)
                        outputArray[i] = inputArray[i + size + shift];
                    else
                        outputArray[i] = inputArray[i + shift];
                }
            }
            else
            {
                outputArray = inputArray;
            }

            var output = new StringBuilder();
            for (int i = 0; i < outputArray.Length; i++)
            {
                output.Append(outputArray[i]).Append(' ');
            }

            Console.WriteLine(output);
        }
    }
}
