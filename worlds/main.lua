local world = require('library.world')
local drawSystem = require('systems.draw')
local animationSystem = require('systems.animation')
local playerSystem = require('systems.player')
local hoverSystem = require('systems.hover')
local debugSystem = require('systems.debug')
local playerEntity = require('entities.player')
local bgEntity = require('entities.bg')
local debugEntity = require('entities.debug')

return function(state)
    local w = world('main')

    w.load = function(arg)
        w.addSystem(drawSystem(state))
        w.addSystem(animationSystem(state))
        w.addSystem(playerSystem(state))
        w.addSystem(debugSystem(state))
        w.addSystem(hoverSystem(state))

        w.addEntity(bgEntity())
        w.addEntity(playerEntity())
        w.addEntity(debugEntity())

        w.updateSystems()
        w.sortEntities()

        for index = 1, #w.systems do
            w.systems[index].load(arg)
        end
    end

    return w
end