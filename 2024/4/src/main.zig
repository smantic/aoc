const std = @import("std");

pub fn main() !void {
    const stdout_file = std.io.getStdOut().writer();
    var bw = std.io.bufferedWriter(stdout_file);
    const stdout = bw.writer();

    var file = try std.fs.cwd().openFile("./input.txt", .{});
    defer file.close();

    const allocator = std.heap.page_allocator;

    var matrixHorizontal = try allocator.alloc([141]u8, 141);
    var matrixVertical = try allocator.alloc([141]u8, 141);
    //var matrixDiagonial = try allocator.alloc([141]u8, 141);

    var sum: i64 = 0;
    var i: usize = 0;
    while (try file.reader().readUntilDelimiterOrEof(&matrixHorizontal[i], '\n')) |_| : (i = i + 1) {
        //try stdout.print("{s}\n", .{line});
    }

    sum = sum + find_xmas(matrixHorizontal);

    while (i <= 140) : (i = i + 1) {
        var j: usize = 0;
        while (j <= 140) : (j = j + 1) {
            matrixVertical[i][j] = matrixHorizontal[j][i];
        }
    }

    sum = sum + find_xmas(matrixVertical);

    try stdout.print("xmas appears {d} times.\n", .{sum});
    try bw.flush(); // don't forget to flush!
}

fn find_xmas(m: [][141]u8) i64 {
    var sum: i64 = 0;
    for (m) |row| {
        var iter = std.mem.window(u8, &row, 4, 1);

        while (iter.next()) |window| {
            if (std.mem.eql(u8, window, "XMAS") or std.mem.eql(u8, window, "SAMX")) {
                sum = sum + 1;
            }
        }
    }

    return sum;
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
