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

    var sum: i64 = 0;
    var i: usize = 0;
    var do: bool = true;
    while (i < tokens.items.len) : (i = i + 1) {
        const tok = tokens.items[i];

        if (tok.Tag == Tag.invalid) {
            continue;
        }

        //std.log.info("token: {any}", .{tok.Tag});

        if (tok.Tag == Tag.do) {
            do = true;
            continue;
        }

        if (tok.Tag == Tag.dont) {
            do = false;
            continue;
        }

        if (tok.Tag == Tag.mul and do) {
            if (i + 2 >= tokens.items.len or tokens.items[i + 1].Tag != Tag.number or tokens.items[i + 2].Tag != Tag.number) {
                continue;
            }

            const x = try std.fmt.parseInt(i64, tokens.items[i + 1].data.?, 10);
            const y = try std.fmt.parseInt(i64, tokens.items[i + 2].data.?, 10);
            sum = sum + (x * y);
            std.log.info("{d} : {d}", .{ x, y });

            i = i + 2;
        }
    }

    try stdout.print("sum: {d} \n", .{sum});
    try bw.flush(); // don't forget to flush!
}

const Tag = enum {
    invalid,
    mul,
    number,
    do,
    dont,
};

const Token = struct { Tag: Tag = undefined, data: ?[]const u8 = undefined };

pub const Tokenizer = struct {
    buffer: []const u8,
    index: usize,
    state: State,

    const State = enum {
        start,
        x,
        y,
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
                                //std.log.debug("{s}", .{self.buffer[self.index .. self.index + 4]});
                                if (self.index + 4 < self.buffer.len and std.mem.eql(u8, self.buffer[self.index .. self.index + 4], "mul(")) {
                                    self.index = self.index + 3;
                                    self.state = State.x;
                                    result.Tag = Tag.mul;
                                    return result;
                                }
                                continue;
                            },
                            'd' => {
                                if (self.index + 7 < self.buffer.len and std.mem.eql(u8, self.buffer[self.index .. self.index + 7], "don't()")) {
                                    //std.log.debug("{s}", .{self.buffer[self.index .. self.index + 7]});
                                    result.Tag = Tag.dont;
                                    return result;
                                }
                                if (self.index + 4 < self.buffer.len and std.mem.eql(u8, self.buffer[self.index .. self.index + 4], "do()")) {
                                    //std.log.debug("{s}", .{self.buffer[self.index .. self.index + 4]});
                                    result.Tag = Tag.do;
                                    return result;
                                }
                            },
                            else => continue,
                        }
                    },
                    .x => {
                        switch (self.buffer[self.index]) {
                            '0'...'9' => continue,
                            ',' => {
                                result.data = self.buffer[start..self.index];
                                result.Tag = Tag.number;
                                self.state = State.y;
                                return result;
                            },
                            else => {
                                result.Tag = Tag.invalid;
                                self.state = State.start;
                                return result;
                            },
                        }
                    },
                    .y => {
                        switch (self.buffer[self.index]) {
                            '0'...'9' => continue,
                            ')' => {
                                result.data = self.buffer[start..self.index];
                                result.Tag = Tag.number;
                                self.state = State.start;
                                return result;
                            },
                            else => {
                                result.Tag = Tag.invalid;
                                self.state = State.start;
                                return result;
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
