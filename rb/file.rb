puts "ruby"

loop do
  Channel.receive do |data|
    puts "got data"
    puts data.inspect
  end
end
