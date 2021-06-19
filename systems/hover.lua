local system = require('library.system')

local mouse = love.mouse

return function(state)
    local s = system({'hover', 'position', 'appearance'})

    s.update = function(dt)
        local mouseX, mouseY = mouse.getPosition()
        local e
        for index = 1, #s.entities do
            e = s.entities[index]
            local width = e.appearance.width
            local height = e.appearance.height
            if e.animation then
                local _, _, w, h = e.animation.getFrame():getViewport()
                width = w
                height = h
            end
            local isWithinX = mouseX >= e.position.x and mouseX <= e.position.x + width
            local isWithinY = mouseY >= e.position.y and mouseY <= e.position.y + height
            local isHovered = isWithinX and isWithinY
            if isHovered ~= e.hover.isHovered then
                print('hover ' .. tostring(isHovered))
            end
            e.hover.isHovered = isHovered
        end
    end

    return s
end