local component = require('library.component')

return function(sheet)
    local c = component('animation')
    c.frame = 1
    c.time = 0
    c.duration = 0
    c.frames = {}
    c.sequences = {}

    for index = 1, #sheet.frames do
        local f = sheet.frames[index]
        c.duration = c.duration + f.duration
        c.frames[#c.frames + 1] = love.graphics.newQuad(
            f.frame.x, f.frame.y,
            f.frame.w, f.frame.h,
            sheet.meta.size.w, sheet.meta.size.h
        )
    end

    for index = 1, #sheet.meta.frameTags do
        local t = sheet.meta.frameTags[index]
        -- Treat tag name ending in '_loop' as meaning it should loop
        -- There's no native way to do this in Aseprite
        local matches = {}
        for match in string.gmatch(t.name, '([^_]+)') do
            matches[#matches + 1] = match
        end
        local shouldLoop = matches[#matches] == 'loop'
        c.sequences[#c.sequences + 1] = {
            name = t.name,
            from = t.from,
            to = t.to,
            shouldLoop = shouldLoop,
            direction = t.direction
        }
    end

    c.getFrame = function()
        return c.frames[c.frame]
    end

    return c
end