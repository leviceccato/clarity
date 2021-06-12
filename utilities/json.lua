local json = require('vendor.json')

return function(name)
    return json.decode(love.filesystem.read(name))
end