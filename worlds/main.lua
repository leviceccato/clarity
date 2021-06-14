local world = require('library.world')
local drawSystem = require('systems.draw')
local animationSystem = require('systems.animation')
local playerSystem = require('systems.player')
local uiSystem = require('systems.ui')
local playerEntity = require('entities.player')
local bgEntity = require('entities.bg')

return function(state)
    local w = world('main')

    w.load = function(arg)
        w.addSystem(drawSystem(state))
        w.addSystem(animationSystem(state))
        w.addSystem(playerSystem(state))
        w.addSystem(uiSystem(state))

        w.addEntity(bgEntity())
        w.addEntity(playerEntity())

        w.updateSystems()
        w.sortEntities()
    end

    return w
end