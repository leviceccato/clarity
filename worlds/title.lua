local world = require('library.world')
local titleBgEntity = require('entities.title-bg')

return function(state)
    local w = world('title')

    w.load = function(arg)
        w.addEntity(titleBgEntity())

        w.updateSystems()
        w.sortEntities()
    end

    return w
end