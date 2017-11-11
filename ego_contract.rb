buf = STDIN.read

buf2 = buf.gsub(/\n\s*if err != nil {\n[^\n\/]+\n\s*}/) do |match|
  " //;" + match.gsub(/\s+/, " ")
end

puts buf2
