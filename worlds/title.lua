local world = require('library.world')
local uiSystem = require('systems.ui')
local titleBgEntity = require('entities.title-bg')

return function(state)
    local w = world('title')

    w.load = function(arg)
        w.addSystem(uiSystem(state))

        w.addEntity(titleBgEntity())

        w.updateSystems()
        w.sortEntities()
    end

    return w
end