const std = @import("std");

pub fn main() !void {
    var file = try std.fs.cwd().openFile("./input.txt", .{});
    defer file.close();

    const stdout_file = std.io.getStdOut().writer();
    var bw = std.io.bufferedWriter(stdout_file);
    const stdout = bw.writer();

    const allocator = std.heap.page_allocator;

    var buffer = try allocator.alloc(u8, 1 << 16);
    const bytes_read = try file.reader().readAll(buffer);

    buffer = buffer[0..bytes_read];

    var tokenizer = Tokenizer.init(buffer);
    var tokens = std.ArrayList(Token).init(allocator);

    while (tokenizer.next()) |token| {
        try tokens.append(token);
    }

    const nums = tokens.items;

    var sum: i64 = 0;
    var i: usize = 0;
    while (i < nums.len) : (i = i + 1) {
        std.log.debug("nums: {s} * {s}", .{ nums[i].x, nums[i].y });
        const x = try std.fmt.parseInt(i64, nums[i].x, 10);
        const y = try std.fmt.parseInt(i64, nums[i].y, 10);
        sum = sum + (x * y);
    }

    try stdout.print("sum: {d} \n", .{sum});
    try bw.flush(); // don't forget to flush!
}

const Tag = enum {
    number,
};

const Token = struct { Tag: Tag = undefined, x: []const u8 = undefined, y: []const u8 = undefined };

pub const Tokenizer = struct {
    buffer: []const u8,
    index: usize,
    state: State,

    const State = enum {
        start,
        mul,
        number,
    };

    pub fn init(buff: []const u8) Tokenizer {
        return .{
            .buffer = buff,
            .index = 0,
            .state = State.start,
        };
    }

    pub fn next(self: *Tokenizer) ?Token {
        var result: Token = .{};
        var start = self.index;

        while (self.index < self.buffer.len) {
            {
                defer self.index = self.index + 1;

                switch (self.state) {
                    .start => {
                        start = self.index;
                        switch (self.buffer[self.index]) {
                            'm' => {
                                self.state = State.mul;
                                continue;
                            },
                            else => continue,
                        }
                    },
                    .mul => {
                        if (self.index + 3 >= self.buffer.len) {
                            self.state = State.start;
                            continue;
                        }
                        if (std.mem.eql(u8, self.buffer[self.index .. self.index + 3], "ul(")) {
                            start = self.index + 3;
                            self.index = self.index + 2;
                            self.state = State.number;
                            continue;
                        } else {
                            self.state = State.start;
                            continue;
                        }
                    },
                    .number => {
                        switch (self.buffer[self.index]) {
                            '0'...'9' => continue,
                            ',' => {
                                result.x = self.buffer[start..self.index];
                                result.Tag = Tag.number;
                                // keep state on num.
                                start = self.index + 1;
                            },
                            ')' => {
                                result.y = self.buffer[start..self.index];

                                self.state = State.start;
                                return result;
                            },
                            else => {
                                self.state = State.start;
                                continue;
                            },
                        }
                    },
                }
            }
        }
        return null;
    }
};

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
