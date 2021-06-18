local system = require('library.system')

local mouse = love.mouse

return function(state)
    local s = system({'hover', 'position', 'appearance'})

    s.update = function(dt)
        local mouseX, mouseY = mouse.getPosition()
        local e
        for index = 1, #s.entities do
            e = s.entities[index]
            local isWithinX = mouseX >= e.position.x and mouseX <= e.position.x + e.appearance.width
            local isWithinY = mouseY >= e.position.y and mouseY <= e.position.y + e.appearance.height
            local isHovered = isWithinX and isWithinY
            if isHovered ~= e.hover.isHovered then
                print('hover')
            end
            e.hover.isHovered = isHovered
        end
    end

    return s
end