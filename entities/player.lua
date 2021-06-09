local entity = require('library.entity')
local appearanceComponent = require('components.appearance')
local positionComponent = require('components.position')

return function()
    local e = entity()
    e.addComponent(appearanceComponent())
    e.addComponent(positionComponent())

    return e
end