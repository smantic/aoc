d = File.read!("input.txt")
s = String.split(d, "\n", trim: true)

{c, _ } = for n <- s, reduce: {-1, 0} do

  acc -> 
    {n, _} = Integer.parse(n) 
    case acc do 
      {acc, x} when n > x ->
        {acc+1, n} 
      {acc, x} when n < x -> {acc, n}
    end 
end

IO.puts(c)
