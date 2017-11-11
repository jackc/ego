STDIN.each_line do |l|
  code, comment = l.split("//;")
  puts code
  puts comment if comment
end
