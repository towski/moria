class Encoding 
  UTF_8 = Encoding.find('UTF-8')
end
#require './env'

require 'rubygems'
require 'bundler'
Encoding.add_alias "binary", "ASCII-8BIT"
Encoding.add_alias "us_ascii", "ASCII-8BIT"
require ::File.expand_path('../minecart/config/environment')

i = 0
loop do
  Channel.receive do |value|
    dwarf_hash = value["Dwarf"]
    civ = Civilization.find_or_initialize_by(:Id => dwarf_hash["CivId"])
    dwarf = Dwarf.find_or_initialize_by(:Id => dwarf_hash["Id"])
    civ.save if civ.new_record?
    dwarf.update_attributes :Name => value["Name"], :CivId => dwarf_hash["CivId"]
    i += 1
    if i % 300 == 0
      puts "processed #{i} dwarves"
    end
  end
end
