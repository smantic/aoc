input = File.read!("input.txt")
  |> String.split("\n")
  |> (fn x -> for i <- x, do: String.split(i, " ", trim: true) end).()
  |> (fn x -> for [ dir, val ] <- x do 
      num = String.to_integer(val) 
      [dir, num]
    end 
  end).()

x = for [dir, val] <- input, reduce: {0, 0} do
  {x, y} -> case dir do 
    "forward" -> {x + val, y} 
    "down" -> {x, y - val}
    "up" -> {x, y + val}
  end
end

IO.inspect(x) 
