local entity = require('library.entity')
local appearance = require('components.appearance')
local animation = require('components.animation')
local position = require('components.position')
local text = require('components.text')
local hover = require('components.hover')
local click = require('components.click')

return function(button)
    local x = button.x or 0
    local y = button.y or 0
    local content = button.content or ''
    local image = button.image
    local sheet = button.sheet
    local align = button.align or 'left'
    local padding = button.padding or 0
    local colour = button.colour or love.c.colFg
    local handler = button.handler

    local e = entity()

    if image then
        e.addComponent(appearance(image))
    end

    if sheet then
        e.addComponent(animation(sheet))
    end

    if handler then
        e.addComponent(click(handler))
    end

    e.addComponent(position(x, y))
    e.addComponent(text(content, align, padding, colour))
    e.addComponent(hover())

    return e
end