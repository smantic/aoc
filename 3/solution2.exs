use Bitwise

defmodule P2 do 
  def fun(input, acc, shift) do 
    freq = for x <- input, reduce: 0 do 
      f -> 
        b = x >>> shift - 1
        cond do 
           (b &&& 1) == 1 -> f + 1
           (b &&& 1) == 0 -> f 
        end 
      end 

    IO.inspect("acc: #{acc}. freq: #{freq}. shift: #{shift}")
    half = length(input) / 2
    cond do 
      freq == 0 -> input
      shift == 0 -> acc
      freq < half -> 
        input = for x <- input, ((x >>> (shift - 1)) &&& 1) == 1, do: x 
        fun(input, acc ||| 1 <<< shift, shift - 1)
      freq >= half -> 
        input = for x <- input, ((x >>> (shift - 1)) &&& 1) == 0, do: x 
        fun(input, acc, shift - 1 )
    end 
  end
end 

input = File.read!("input.txt")
  |> String.split("\n", trim: true)
  |> (fn all -> 
    for bits <- all do 
      String.to_integer(bits, 2)
    end 
  end).()

P2.fun(input, 0, 12)
|> IO.inspect
