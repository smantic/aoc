defmodule PartTwo do
  def fun(elems, prev) do
    [ _ | slide] = elems
    case elems do
      [a, b, c | _ ] when a + b + c  > prev -> 1 + fun(slide, a + b + c) 
      [a, b, c | _ ] when a + b + c  <= prev -> fun(slide, a + b + c) 
      _ -> -1
    end 
  end
end

File.read!("input.txt")
  |> String.split("\n", trim: true)
  |> Enum.map(&String.to_integer/1)
  |> PartTwo.fun(0)
  |> IO.puts


