using System;
using System.Collections.Generic;

namespace JSZYCER
{
    class Program
    {
        static void Main()
        {
            int input;
            IDictionary<char, char> dict = new Dictionary<char, char>
            {
                {'a', 'd'}, {'A', 'D'},
                {'b', 'e'}, {'B', 'E'},
                {'c', 'f'}, {'C', 'F'},
                {'d', 'g'}, {'D', 'G'},
                {'e', 'h'}, {'E', 'H'},
                {'f', 'i'}, {'F', 'I'},
                {'g', 'j'}, {'G', 'J'},
                {'h', 'k'}, {'H', 'K'},
                {'i', 'l'}, {'I', 'L'},
                {'j', 'm'}, {'J', 'M'},
                {'k', 'n'}, {'K', 'N'},
                {'l', 'o'}, {'L', 'O'},
                {'m', 'p'}, {'M', 'P'},
                {'n', 'q'}, {'N', 'Q'},
                {'o', 'r'}, {'O', 'R'},
                {'p', 's'}, {'P', 'S'},
                {'q', 't'}, {'Q', 'T'},
                {'r', 'u'}, {'R', 'U'},
                {'s', 'v'}, {'S', 'V'},
                {'t', 'w'}, {'T', 'W'},
                {'u', 'x'}, {'U', 'X'},
                {'v', 'y'}, {'V', 'Y'},
                {'w', 'z'}, {'W', 'Z'},
                {'x', 'a'}, {'X', 'A'},
                {'y', 'b'}, {'Y', 'B'},
                {'z', 'c'}, {'Z', 'C'}
            };

            while ((input = Console.Read()) != -1)
            {
                if (dict.ContainsKey((char)input))
                {
                    Console.Write(dict[(char)input]);
                }
                else
                {
                    Console.Write((char)input);
                }
            }
        }
    }
}
