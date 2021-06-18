local system = require('library.system')

local floor = math.floor

return function(state)
    local s = system({'animation'})

    s.update = function(dt)
        for index = 1, #s.entities do
            local e = s.entities[index]
            e.animation.time = e.animation.time + (dt * 1000)
            if e.animation.time >= e.animation.duration then
                e.animation.time = e.animation.time - e.animation.duration
            end
            local sequence = e.animation.sequences[e.animation.sequence]
            local length = sequence.to - sequence.from
            e.animation.frame = floor(e.animation.time / e.animation.duration * length) + 1 + sequence.from
        end
    end

    return s
end