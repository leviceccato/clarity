local world = require('library.world')
local drawSystem = require('systems.draw')
local animationSystem = require('systems.animation')
local debugSystem = require('systems.debug')
local hoverSystem = require('systems.hover')
local titleBgEntity = require('entities.title-bg')

return function(state)
    local w = world('title')

    w.load = function(arg)
        w.addSystem(drawSystem(state))
        w.addSystem(animationSystem(state))
        w.addSystem(debugSystem(state))
        w.addSystem(hoverSystem(state))

        w.addEntity(titleBgEntity())

        w.updateSystems()
        w.sortEntities()

        for index = 1, #w.systems do
            w.systems[index].load(arg)
        end
    end

    return w
end