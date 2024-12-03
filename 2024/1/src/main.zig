const std = @import("std");

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    var l1 = std.ArrayList(i64).init(allocator);
    var l2 = std.ArrayList(i64).init(allocator);
    defer l1.deinit();
    defer l2.deinit();

    var file = try std.fs.cwd().openFile("./input.txt", .{});
    defer file.close();

    var buf: [100 * 1000]u8 = undefined;
    const read_bytes = try file.reader().readAll(&buf);
    buf[read_bytes] = 0;

    var iter = std.mem.tokenizeAny(u8, &buf, &std.ascii.whitespace);

    var count: i32 = 0;
    while (iter.next()) |part| : (count += 1) {
        const i = std.fmt.parseInt(i64, part, 10) catch break;
        if (@mod(count, 2) == 0) {
            try l1.append(i);
        } else {
            try l2.append(i);
        }
    }

    std.mem.sort(i64, l1.items, {}, std.sort.asc(i64));
    std.mem.sort(i64, l2.items, {}, std.sort.asc(i64));

    var i: usize = 0;
    var sum: u64 = 0;
    while (i < l1.items.len) : (i += 1) {
        sum = sum + @abs(l1.items[i] - l2.items[i]);
    }

    std.log.info("total distance: {d}\n", .{sum});
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
