const std = @import("std");

pub fn main() !void {
    // Prints to stderr (it's a shortcut based on `std.io.getStdErr()`)
    std.debug.print("All your {s} are belong to us.\n", .{"codebase"});

    var file = try std.fs.cwd().openFile("./input.txt", .{});
    defer file.close();

    // stdout is for the actual output of your application, for example if you
    // are implementing gzip, then only the compressed bytes should be sent to
    // stdout, not any debugging messages.
    const stdout_file = std.io.getStdOut().writer();
    var bw = std.io.bufferedWriter(stdout_file);
    const stdout = bw.writer();

    var buf: [256]u8 = undefined;

    var sum: i32 = 0;
    while (true) {
        const line = file.reader().readUntilDelimiter(&buf, '\n') catch break;

        var iter = std.mem.tokenizeAny(u8, line, &std.ascii.whitespace);

        var prev = std.fmt.parseInt(i64, iter.next().?, 10) catch break;
        var isSafe = true;
        var isAscending = false;
        var isFirstItem = true;

        while (iter.next()) |part| {
            const i = std.fmt.parseInt(i64, part, 10) catch break;
            defer prev = i;

            try stdout.print("i: {d}\n", .{i});
            const difference = prev - i;
            if ((isAscending != (difference < 0)) and !isFirstItem) {
                isSafe = false;
                break;
            } else {
                isAscending = (difference < 0);
                isFirstItem = false;
            }

            const abs_difference = @abs(difference);
            if (abs_difference == 0 or abs_difference > 3) {
                isSafe = false;
                break;
            }
        }

        if (isSafe) {
            sum = sum + 1;
        }
    }

    try stdout.print("safe tests: {d}", .{sum});
    try bw.flush(); // don't forget to flush!
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
