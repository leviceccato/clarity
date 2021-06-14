local system = require('library.system')

local keyboard = love.keyboard

return function(state)
    local s = system({'position', 'controls'})

    s.update = function()
        local e
        for index = 1, #s.entities do
            e = s.entities[index]
            if keyboard.isDown('w') then
                e.position.y = e.position.y - 2
            end
            if keyboard.isDown('a') then
                e.position.x = e.position.x - 2
            end
            if keyboard.isDown('s') then
                e.position.y = e.position.y + 2
            end
            if keyboard.isDown('d') then
                e.position.x = e.position.x + 2
            end
        end
    end

    return s
end