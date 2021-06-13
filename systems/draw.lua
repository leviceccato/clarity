local system = require('library.system')

local lgDraw = love.graphics.draw

return function()
    local s = system({'appearance', 'position'})

    s.draw = function()
        for index = 1, #s.entities do
            local e = s.entities[index]
            lgDraw(e.appearance.image, e.position.x, e.position.y)
        end
    end

    return s
end