local world = require('library.world')
local drawSystem = require('systems.draw')
local playerSystem = require('systems.player')
local uiSystem = require('systems.ui')
local playerEntity = require('entities.player')
local bgEntity = require('entities.bg')

return function()
    local w = world('main')

    w.load = function(arg)
        w.addSystem(drawSystem())
        w.addSystem(playerSystem())
        w.addSystem(uiSystem())

        w.addEntity(bgEntity())
        w.addEntity(playerEntity())

        w.updateSystems()
        w.sortEntities()
    end

    return w
end