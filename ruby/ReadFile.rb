class ReadFile

	def initialize
		@file = File.open( "../data.json" ) 
	end

	def readFile
		t1 			= Time.now
		filesize 	= @file.stat.size
		chunck 		= 64*1024
		position 	= 0
		n_chunck	= 0

		while position <= filesize
			buffer = @file.read( chunck )
			position += chunck
			@file.seek(position, File::SEEK_SET);
			n_chunck += 1
		end

		t2			= Time.now
		puts "Took: " + time_diff_milli(t1, t2).to_s + " ms - chunck: " + n_chunck.to_s
	end

	def time_diff_milli(start, finish)
   		(finish - start) * 1000.0
	end
end

if __FILE__ == $0
	read = ReadFile.new
	read.readFile
end