local system = require('library.system')

return function()
    local s = system()
    s.components = { 'position', 'controls' }

    s.update = function()
        local e
        for index = 1, #s.entities do
            e = s.entities[index]
            if love.keyboard.isDown('w') then
                e.position.y = e.position.y - 2
            end
            if love.keyboard.isDown('a') then
                e.position.x = e.position.x - 2
            end
            if love.keyboard.isDown('s') then
                e.position.y = e.position.y + 2
            end
            if love.keyboard.isDown('d') then
                e.position.x = e.position.x + 2
            end
        end
    end

    return s
end