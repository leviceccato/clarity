local system = require('library.system')

return function()
    local s = system()
    s.components = { 'appearance', 'position' }

    s.update = function()
        local entity
        for index = 1, #s.entities do
            entity = s.entities[index]
        end
    end

    return s
end