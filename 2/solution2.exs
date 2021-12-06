input = File.read!("input.txt")
  |> String.split("\n")
  |> (fn x -> for i <- x, do: String.split(i, " ", trim: true) end).()
  |> (fn x -> for [ dir, val ] <- x do 
      num = String.to_integer(val) 
      [dir, num]
    end 
  end).()

x = for [dir, val] <- input, reduce: %{x: 0, y: 0, aim: 0} do
  %{x: x, y: y, aim: z} -> case dir do 
    "up" -> %{ x: x, y: y, aim: z - val }
    "down" -> %{ x: x, y: y, aim: z + val }
    "forward" -> %{ x: x + val, y: y + (val * z), aim: z } 
  end
end

IO.inspect(x) 
