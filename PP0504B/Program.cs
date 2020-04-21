using System;
using System.Text;

namespace PP0504B
{
    class Program
    {
        static void Main()
        {
            short rowCount = Convert.ToInt16(Console.ReadLine());

            while (rowCount-- > 0)
            {
                StringBuilder output = new StringBuilder();
                string[] strings = Console.ReadLine().Split(' ');
                int length;

                if (strings[0].Length < strings[1].Length)
                {
                    length = strings[0].Length;
                }
                else
                {
                    length = strings[1].Length;
                }

                for (int i = 0; i < length; i++)
                {
                    output.Append(strings[0].Substring(i, 1));
                    output.Append(strings[1].Substring(i, 1));
                }

                Console.WriteLine(output);
            }
        }
    }
}
