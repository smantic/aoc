input = File.read!("input.txt")
  |> String.split("\n", trim: true)
  |> (fn x -> for i <- x, do: String.split(i, "", trim: true) end).()
  |> (fn all ->
    for list <- all do
      for bit <- list do 
        case bit do 
          "0" -> -1 
          "1" -> 1
        end
      end
    end 
  end).()

x  = for bits <- input, reduce: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]  do 
  acc -> Enum.zip_with(acc, bits, fn x, y -> x + y end)
end

gamma = for count <- x do  
  cond do 
    count > 0 -> 1
    count < 0 -> 0
  end 
end 

eps = for count <- x do  
  cond do 
    count > 0 -> 0
    count < 0 -> 1
  end 
end 


IO.inspect(gamma)
IO.inspect(eps)
